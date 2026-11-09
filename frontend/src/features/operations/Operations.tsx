import {FormEvent, useEffect, useState} from 'react';
import {api} from '../../auth/session';
import {Layout} from '../dashboard/Layout';
import {gql} from '../graphql/client';

type QueueItem = {
  question: {
    id: string;
    body: string;
    priorityClass: string;
    priorityScore: number;
    lifecycleStatus: string;
    matchedUrgencyWords: string[];
    priorityDictionaryVersion: string;
    createdAt: string;
    lastCategorizedAt: string;
  };
  order: {id: string; status: string; total: number};
  seller: {id: string; name: string; email: string};
  buyer: {name: string};
  waitingDays: number;
};

type Filters = {seller: string; buyer: string; status: string; dateFrom: string; dateTo: string; priorityClass: string};
const emptyFilters: Filters = {seller: '', buyer: '', status: '', dateFrom: '', dateTo: '', priorityClass: ''};

export function Operations() {
  const [items, setItems] = useState<QueueItem[]>([]);
  const [filters, setFilters] = useState<Filters>(emptyFilters);
  const [applied, setApplied] = useState<Filters>(emptyFilters);
  const [feedback, setFeedback] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(true);
  const [reviewing, setReviewing] = useState(false);

  async function load() {
    setLoading(true);
    setError('');
    try {
      const data = await gql<{operationsPriorityQueue: {questions: QueueItem[]}}>(
        `query OperationsPriorityQueue($filters:PriorityFilters){
          operationsPriorityQueue(filters:$filters){
            questions {
              question {
                id body priorityClass priorityScore lifecycleStatus matchedUrgencyWords
                priorityDictionaryVersion createdAt lastCategorizedAt
              }
              order { id status total }
              seller { id name email }
              buyer { name }
              waitingDays
            }
          }
        }`,
        {filters: applied},
      );
      setItems(data.operationsPriorityQueue.questions);
    } catch (reason) {
      setError((reason as Error).message);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    void load();
  }, [applied]);

  async function review() {
    setReviewing(true);
    setError('');
    try {
      const data = await api<{review: {status: string; recategorized: number; notificationBatchesLogged: number; executionId: string}}>(
        '/api/operations/priority-review',
        {method: 'POST', body: '{}'},
      );
      setFeedback(`${data.review.status} · ${data.review.recategorized} recategorizadas · ${data.review.notificationBatchesLogged} lotes enviados · ejecución ${data.review.executionId}`);
      await load();
    } catch (reason) {
      setError((reason as Error).message);
    } finally {
      setReviewing(false);
    }
  }

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
          <span className="eyebrow">OPERACIONES</span>
          <h2>Cola de prioridad</h2>
          <p>Preguntas abiertas que necesitan atención.</p>
        </div>
        <button onClick={review} disabled={reviewing}>{reviewing ? 'Revisando…' : 'Ejecutar revisión'}</button>
      </div>
      {feedback && <div className="success">{feedback}</div>}
      {error && <div className="error">{error}</div>}
      <form className="filters" onSubmit={submitFilters}>
        <input aria-label="Vendedor" placeholder="Vendedor, correo o ID" value={filters.seller} onChange={event => setFilters({...filters, seller: event.target.value})} />
        <input aria-label="Comprador" placeholder="Comprador" value={filters.buyer} onChange={event => setFilters({...filters, buyer: event.target.value})} />
        <select aria-label="Estado de orden" value={filters.status} onChange={event => setFilters({...filters, status: event.target.value})}>
          <option value="">Todos los estados</option>
          {['new', 'processing', 'shipped', 'delivered', 'cancelled'].map(value => <option key={value}>{value}</option>)}
        </select>
        <select aria-label="Prioridad" value={filters.priorityClass} onChange={event => setFilters({...filters, priorityClass: event.target.value})}>
          <option value="">Todas las prioridades</option>
          {['Critical', 'High Priority', 'Medium', 'Low'].map(value => <option key={value}>{value}</option>)}
        </select>
        <input aria-label="Desde" type="date" value={filters.dateFrom} onChange={event => setFilters({...filters, dateFrom: event.target.value})} />
        <input aria-label="Hasta" type="date" value={filters.dateTo} onChange={event => setFilters({...filters, dateTo: event.target.value})} />
        <button type="submit">Aplicar filtros</button>
        <button type="button" className="ghost" onClick={clearFilters}>Limpiar</button>
      </form>
      {loading ? <div className="skeleton">Cargando cola…</div> : (
        <div className="queue">
          {items.length ? items.map(item => (
            <article className="card queue-item" key={item.question.id}>
              <div className="score">{item.question.priorityScore}</div>
              <div>
                <span className="pill priority">{item.question.priorityClass}</span>
                <h3>{item.question.body}</h3>
                <p>{item.seller.name} ({item.seller.id}) · {item.buyer.name} · {item.order.id}</p>
                <small>{item.waitingDays} días esperando · Orden ${item.order.total.toFixed(2)}</small>
                <small>Creada: {new Date(item.question.createdAt).toLocaleString()} · Categorizada: {new Date(item.question.lastCategorizedAt).toLocaleString()}</small>
                <small>Diccionario: {item.question.priorityDictionaryVersion} · Coincidencias: {item.question.matchedUrgencyWords.join(', ') || 'ninguna'}</small>
              </div>
            </article>
          )) : <div className="empty">No hay preguntas sin resolver.</div>}
        </div>
      )}
    </Layout>
  );
}
