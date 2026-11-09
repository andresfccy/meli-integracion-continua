import {FormEvent, useEffect, useState} from 'react';
import {api, getSession} from '../../auth/session';
import {Layout} from '../dashboard/Layout';
import {gql} from '../graphql/client';

type Question = {
  id: string;
  buyerId: string;
  orderItemId?: string;
  body: string;
  status: string;
  lifecycleStatus: string;
  createdAt: string;
  reply?: {body: string};
};
type Order = {
  id: string;
  status: string;
  orderedAt: string;
  total: number;
  seller: {id: string; name: string; email: string};
  items: {id: string; productName: string; quantity: number; unitPrice: number}[];
  questions: Question[];
};
type Filters = {status: string; seller: string; dateFrom: string; dateTo: string};
const emptyFilters: Filters = {status: '', seller: '', dateFrom: '', dateTo: ''};
const orderFields = `
  id status orderedAt total
  seller { id name email }
  items { id productName quantity unitPrice }
  questions { id buyerId orderItemId body status lifecycleStatus createdAt reply { body } }
`;

export function Buyer() {
  const session = getSession()!;
  const orderId = location.pathname.split('/')[3];
  const [orders, setOrders] = useState<Order[]>([]);
  const [order, setOrder] = useState<Order | null>(null);
  const [filters, setFilters] = useState<Filters>(emptyFilters);
  const [applied, setApplied] = useState<Filters>(emptyFilters);
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    void (async () => {
      setLoading(true);
      setError('');
      try {
        if (orderId) {
          const data = await gql<{buyerOrderDetail: Order}>(
            `query BuyerOrderDetail($buyerId:String!,$orderId:String!){
              buyerOrderDetail(buyerId:$buyerId,orderId:$orderId){ ${orderFields} }
            }`,
            {buyerId: session.user.buyerId!, orderId},
          );
          setOrder(data.buyerOrderDetail);
        } else {
          const data = await gql<{buyerDashboard: {orders: Order[]}}>(
            `query BuyerDashboard($buyerId:String!,$filters:OrderFilters){
              buyerDashboard(buyerId:$buyerId,filters:$filters){ orders { ${orderFields} } }
            }`,
            {buyerId: session.user.buyerId!, filters: applied},
          );
          setOrders(data.buyerDashboard.orders);
        }
      } catch (reason) {
        setError((reason as Error).message);
      } finally {
        setLoading(false);
      }
    })();
  }, [orderId, applied, session.user.buyerId]);

  function submitFilters(event: FormEvent) {
    event.preventDefault();
    setApplied({...filters});
  }

  return (
    <Layout>
      <div className="title">
        <div>
          <span className="eyebrow">MIS COMPRAS</span>
          <h2>{orderId ? 'Detalle de compra' : 'Tus órdenes'}</h2>
        </div>
        {orderId && <a href="/buyer/orders">← Volver</a>}
      </div>
      {error ? <div className="error">{error}</div> : loading ? <div className="skeleton">Cargando…</div> : orderId ? (
        order ? <Detail order={order} buyerId={session.user.buyerId!} /> : <div className="empty">Orden no encontrada.</div>
      ) : (
        <>
          <form className="filters" onSubmit={submitFilters}>
            <input aria-label="Vendedor" placeholder="Vendedor, correo o ID" value={filters.seller} onChange={event => setFilters({...filters, seller: event.target.value})} />
            <select aria-label="Estado" value={filters.status} onChange={event => setFilters({...filters, status: event.target.value})}>
              <option value="">Todos los estados</option>
              {['new', 'processing', 'shipped', 'delivered', 'cancelled'].map(value => <option key={value}>{value}</option>)}
            </select>
            <input aria-label="Desde" type="date" value={filters.dateFrom} onChange={event => setFilters({...filters, dateFrom: event.target.value})} />
            <input aria-label="Hasta" type="date" value={filters.dateTo} onChange={event => setFilters({...filters, dateTo: event.target.value})} />
            <button type="submit">Aplicar filtros</button>
            <button type="button" className="ghost" onClick={() => {setFilters(emptyFilters); setApplied(emptyFilters);}}>Limpiar</button>
          </form>
          <div className="grid">
            {orders.length ? orders.map(value => (
              <a className="card order" href={`/buyer/orders/${value.id}`} key={value.id}>
                <span className={`pill ${value.status}`}>{value.status}</span>
                <h3>{value.items[0].productName}</h3>
                <p>{value.seller.name} · {new Date(value.orderedAt).toLocaleDateString()}</p>
                <strong>${value.total.toFixed(2)}</strong>
                {value.questions.length > 0 && <small>{value.questions[0].lifecycleStatus.replace('_', ' ')}</small>}
              </a>
            )) : <div className="empty">No hay órdenes que coincidan con los filtros.</div>}
          </div>
        </>
      )}
    </Layout>
  );
}

function Detail({order, buyerId}: {order: Order; buyerId: string}) {
  const [questions, setQuestions] = useState(order.questions);
  const [body, setBody] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [busy, setBusy] = useState(false);
  const hasActiveQuestion = questions.some(question => question.status === 'open');

  async function ask(event: FormEvent) {
    event.preventDefault();
    setBusy(true);
    setError('');
    try {
      const result = await api<{question: Question}>(`/api/buyers/${buyerId}/orders/${order.id}/questions`, {
        method: 'POST',
        body: JSON.stringify({body}),
      });
      setQuestions(current => [...current, result.question]);
      setBody('');
      setSuccess('Pregunta enviada.');
    } catch (reason) {
      setError((reason as Error).message);
    } finally {
      setBusy(false);
    }
  }

  async function approve(questionId: string) {
    setBusy(true);
    setError('');
    try {
      const result = await api<{question: Question}>(`/api/buyers/${buyerId}/questions/${questionId}/approve-answer`, {method: 'POST', body: '{}'});
      setQuestions(current => current.map(question => question.id === questionId ? result.question : question));
      setSuccess('Respuesta aprobada y pregunta cerrada.');
    } catch (reason) {
      setError((reason as Error).message);
    } finally {
      setBusy(false);
    }
  }

  return (
    <div className="columns">
      <section className="card">
        <span className={`pill ${order.status}`}>{order.status}</span>
        <h3>{order.id}</h3>
        <p>{order.seller.name} · {order.seller.email}</p>
        {order.items.map(item => (
          <div className="item" key={item.id}>
            <span>{item.quantity} × {item.productName}</span>
            <b>${(item.quantity * item.unitPrice).toFixed(2)}</b>
          </div>
        ))}
        <div className="total"><span>Total</span><b>${order.total.toFixed(2)}</b></div>
      </section>
      <section>
        <h3>Preguntas</h3>
        {error && <div className="error">{error}</div>}
        {success && <div className="success">{success}</div>}
        {questions.map(question => (
          <article className="card question" key={question.id}>
            <span className="pill lifecycle">{question.lifecycleStatus.replace('_', ' ')}</span>
            <small>{new Date(question.createdAt).toLocaleString()} · {order.items.find(item => item.id === question.orderItemId)?.productName ?? 'Pregunta general de la orden'}</small>
            <p>{question.body}</p>
            {question.reply && (
              <>
                <blockquote>{question.reply.body}</blockquote>
                {question.lifecycleStatus === 'seller_answered' && <button disabled={busy} onClick={() => approve(question.id)}>{busy ? 'Aprobando…' : 'Aprobar respuesta'}</button>}
              </>
            )}
          </article>
        ))}
        {!hasActiveQuestion && (
          <form className="card question" onSubmit={ask}>
            <label>¿Necesitas ayuda con esta compra?
              <textarea value={body} onChange={event => setBody(event.target.value)} placeholder="Escribe tu pregunta…" required />
            </label>
            <button disabled={busy}>{busy ? 'Enviando…' : 'Enviar pregunta'}</button>
          </form>
        )}
      </section>
    </div>
  );
}
