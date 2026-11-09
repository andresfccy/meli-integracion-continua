import type {ReactNode} from 'react';
import {clearSession, getSession} from '../../auth/session';

export function Layout({children}: {children: ReactNode}) {
  const session = getSession()!;
  return (
    <>
      <header>
        <div><span className="eyebrow">MERCADO</span><h1>Centro de gestión</h1></div>
        <div className="identity">
          <span>{session.user.displayName}</span>
          <button className="ghost" onClick={() => {clearSession(); location.href = '/login';}}>Cerrar sesión</button>
        </div>
      </header>
      <main>{children}</main>
    </>
  );
}
