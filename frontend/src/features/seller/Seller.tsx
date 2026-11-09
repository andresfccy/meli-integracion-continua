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
  priorityClass: string;
  priorityScore: number;
  createdAt: string;
  reply?: {body: string};
};

type Order = {
  id: string;
  status: string;
  orderedAt: string;
  total: number;
  buyer: {name: string; email: string};
  items: {id: string; productName: string; quantity: number; unitPrice: number}[];
  questions: Question[];
};

type Filters = {status: string; buyer: string; dateFrom: string; dateTo: string};
const emptyFilters: Filters = {status: '', buyer: '', dateFrom: '', dateTo: ''};
const orderFields = `
  id status orderedAt total
  buyer { name email }
  items { id productName quantity unitPrice }
  questions {
    id buyerId orderItemId body status lifecycleStatus priorityClass priorityScore createdAt
    reply { body }
  }
`;

export function Seller() {
  const session = getSession()!;
  const detail = location.pathname.split('/')[3];
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
        if (detail) {
          const data = await gql<{sellerOrderDetail: Order}>(
            `query SellerOrderDetail($sellerId:String!,$orderId:String!){
              sellerOrderDetail(sellerId:$sellerId,orderId:$orderId){ ${orderFields} }
            }`,
            {sellerId: session.user.sellerId!, orderId: detail},
          );
          setOrder(data.sellerOrderDetail);
        } else {
          const data = await gql<{sellerDashboard: {orders: Order[]}}>(
            `query SellerDashboard($sellerId:String!,$filters:OrderFilters){
              sellerDashboard(sellerId:$sellerId,filters:$filters){ orders { ${orderFields} } }
            }`,
            {sellerId: session.user.sellerId!, filters: applied},
          );
          setOrders(data.sellerDashboard.orders);
        }
      } catch (reason) {
        setError((reason as Error).message);
      } finally {
        setLoading(false);
      }
    })();
  }, [detail, applied, session.user.sellerId]);

  function submitFilters(event: FormEvent) {
    event.preventDefault();
    setApplied({...filters});
  }

  function clearFilters() {
    setFilters(emptyFilters);
    setApplied(emptyFilters);
  }

  return (
    <Layout>
      <div className="title">
        <div>
          <span className="eyebrow">VENTAS</span>
          <h2>{detail ? 'Detalle de la orden' : 'Órdenes recientes'}</h2>
        </div>
        {detail && <a href="/seller/orders">← Volver</a>}
      </div>
      {error ? (
        <div className="error">{error}</div>
      ) : loading ? (
        <Skeleton />
      ) : detail ? (
        order ? <OrderDetail order={order} sellerId={session.user.sellerId!} /> : <div className="empty">Orden no encontrada.</div>
      ) : (
        <>
          <form className="filters" onSubmit={submitFilters}>
            <input aria-label="Buscar comprador" placeholder="Nombre o correo del comprador" value={filters.buyer} onChange={event => setFilters({...filters, buyer: event.target.value})} />
            <select aria-label="Estado" value={filters.status} onChange={event => setFilters({...filters, status: event.target.value})}>
              <option value="">Todos los estados</option>
              {['new', 'processing', 'shipped', 'delivered', 'cancelled'].map(value => <option key={value}>{value}</option>)}
            </select>
            <input aria-label="Desde" type="date" value={filters.dateFrom} onChange={event => setFilters({...filters, dateFrom: event.target.value})} />
            <input aria-label="Hasta" type="date" value={filters.dateTo} onChange={event => setFilters({...filters, dateTo: event.target.value})} />
            <button type="submit">Aplicar filtros</button>
            <button type="button" className="ghost" onClick={clearFilters}>Limpiar</button>
          </form>
          <OrderList orders={orders} />
        </>
      )}
    </Layout>
  );
}

function OrderList({orders}: {orders: Order[]}) {
  return (
    <div className="grid">
      {orders.length ? orders.map(order => (
        <a className="card order" href={`/seller/orders/${order.id}`} key={order.id}>
          <div><b>{order.id}</b><span className={`pill ${order.status}`}>{order.status}</span></div>
          <h3>{order.items[0]?.productName}</h3>
          <p>{order.buyer.name} · {new Date(order.orderedAt).toLocaleDateString()} · {order.items.length} producto(s)</p>
          <strong>${order.total.toFixed(2)}</strong>
          <small>{order.questions.length ? `${order.questions.length} pregunta(s)` : 'Sin preguntas'}</small>
        </a>
      )) : <div className="empty">No hay órdenes que coincidan con los filtros.</div>}
    </div>
  );
}

function OrderDetail({order, sellerId}: {order: Order; sellerId: string}) {
  return (
    <div className="columns">
      <section className="card">
        <span className={`pill ${order.status}`}>{order.status}</span>
        <h3>{order.id}</h3>
        <p>{order.buyer.name} · {order.buyer.email}</p>
        <p>{new Date(order.orderedAt).toLocaleString()}</p>
        {order.items.map(item => (
          <div className="item" key={item.id}>
            <span>{item.quantity} × {item.productName}</span>
            <b>${(item.quantity * item.unitPrice).toFixed(2)}</b>
          </div>
        ))}
        <div className="total"><span>Total</span><b>${order.total.toFixed(2)}</b></div>
      </section>
      <section>
        <h3>Preguntas del comprador</h3>
        {order.questions.length
          ? order.questions.map(question => (
            <QuestionCard
              key={question.id}
              question={question}
              sellerId={sellerId}
              buyerName={order.buyer.name}
              productName={order.items.find(item => item.id === question.orderItemId)?.productName}
            />
          ))
          : <div className="empty">Esta orden todavía no tiene preguntas.</div>}
      </section>
    </div>
  );
}

function QuestionCard({question, sellerId, buyerName, productName}: {
  question: Question;
  sellerId: string;
  buyerName: string;
  productName?: string;
}) {
  const [current, setCurrent] = useState(question);
  const [body, setBody] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [busy, setBusy] = useState(false);

  async function submit(event: FormEvent) {
    event.preventDefault();
    setBusy(true);
    setError('');
    try {
      const result = await api<{question: Question}>(`/api/sellers/${sellerId}/questions/${current.id}/reply`, {
        method: 'POST',
        body: JSON.stringify({body}),
      });
      setCurrent(result.question);
      setBody('');
      setSuccess('Respuesta guardada.');
    } catch (reason) {
      setError((reason as Error).message);
    } finally {
      setBusy(false);
    }
  }

  return (
    <article className="card question">
      <div>
        <span className="pill priority">{current.priorityClass} · {current.priorityScore}</span>
        <span className="pill lifecycle">{current.lifecycleStatus.replace('_', ' ')}</span>
      </div>
      <small>{buyerName} · {new Date(current.createdAt).toLocaleString()} · {productName ?? 'Pregunta general de la orden'}</small>
      <p>{current.body}</p>
      {success && <div className="success">{success}</div>}
      {current.reply ? <blockquote>{current.reply.body}</blockquote> : (
        <form onSubmit={submit}>
          <textarea aria-label="Respuesta" placeholder="Escribe una respuesta clara…" value={body} onChange={event => setBody(event.target.value)} required />
          {error && <div className="error">{error}</div>}
          <button disabled={busy}>{busy ? 'Enviando…' : 'Responder'}</button>
        </form>
      )}
    </article>
  );
}

function Skeleton() {
  return <div className="skeleton">Cargando información…</div>;
}
