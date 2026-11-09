import {FormEvent, useState} from 'react';
import {api, Session, setSession} from '../../auth/session';

const demos = [
  ['Vendedor', 'seller@example.com', 'seller-demo-password'],
  ['Comprador', 'buyer@example.com', 'buyer-demo-password'],
  ['Operaciones', 'ops@example.com', 'ops-demo-password'],
];

export function Login() {
  const [email, setEmail] = useState('seller@example.com');
  const [password, setPassword] = useState('seller-demo-password');
  const [error, setError] = useState('');
  const [busy, setBusy] = useState(false);

  async function submit(event: FormEvent) {
    event.preventDefault();
    setBusy(true);
    setError('');
    try {
      const session = await api<Session>('/api/auth/login', {
        method: 'POST',
        body: JSON.stringify({email, password}),
      });
      setSession(session);
      location.href = '/';
    } catch (reason) {
      setError((reason as Error).message);
    } finally {
      setBusy(false);
    }
  }

  return (
    <main className="login">
      <section className="login-copy">
        <span className="eyebrow">SELLER DASHBOARD</span>
        <h1>Todo lo importante,<br />en un solo lugar.</h1>
        <p>Órdenes, conversaciones y prioridades para que cada compra avance.</p>
      </section>
      <section className="card login-card">
        <h2>Ingresar</h2>
        <p>Usa una cuenta de demostración.</p>
        <form onSubmit={submit}>
          <label>Correo
            <input value={email} onChange={event => setEmail(event.target.value)} type="email" required />
          </label>
          <label>Contraseña
            <input value={password} onChange={event => setPassword(event.target.value)} type="password" required />
          </label>
          {error && <div className="error">{error}</div>}
          <button disabled={busy}>{busy ? 'Ingresando…' : 'Continuar'}</button>
        </form>
        <div className="demos">
          {demos.map(demo => (
            <button className="demo" key={demo[0]} onClick={() => {setEmail(demo[1]); setPassword(demo[2]);}}>
              <b>{demo[0]}</b><small>{demo[1]}</small>
            </button>
          ))}
        </div>
      </section>
    </main>
  );
}
