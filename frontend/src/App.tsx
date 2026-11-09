import {getSession} from './auth/session';
import {defaultPath} from './auth/routing';
import {Login} from './features/auth/Login';
import {Buyer} from './features/buyer/Buyer';
import {Operations} from './features/operations/Operations';
import {Seller} from './features/seller/Seller';

export default function App() {
  const session = getSession();
  const path = location.pathname;

  if (path === '/login') {
    if (session) {
      location.replace(defaultPath(session.user.roles));
      return null;
    }
    return <Login />;
  }
  if (!session) {
    location.replace('/login');
    return null;
  }
  if (path === '/') {
    location.replace(defaultPath(session.user.roles));
    return null;
  }
  if (path.startsWith('/seller/orders') && session.user.roles.includes('seller')) return <Seller />;
  if (path.startsWith('/buyer/orders') && session.user.roles.includes('buyer')) return <Buyer />;
  if (path === '/operations/questions' && session.user.roles.includes('operations')) return <Operations />;

  location.replace(defaultPath(session.user.roles));
  return null;
}
