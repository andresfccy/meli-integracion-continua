export type User = {
  id: string;
  email: string;
  displayName: string;
  roles: string[];
  sellerId?: string;
  buyerId?: string;
  opsTeamMemberId?: string;
};

export type Session = {user: User; accessToken: string};

const key = 'seller-dashboard-session';

function isSession(value: unknown): value is Session {
  if (!value || typeof value !== 'object') return false;
  const candidate = value as Partial<Session>;
  return typeof candidate.accessToken === 'string'
    && !!candidate.user
    && typeof candidate.user.id === 'string'
    && Array.isArray(candidate.user.roles);
}

export function getSession(): Session | null {
  try {
    const value: unknown = JSON.parse(sessionStorage.getItem(key) || 'null');
    return isSession(value) ? value : null;
  } catch {
    return null;
  }
}

export const setSession = (session: Session) => sessionStorage.setItem(key, JSON.stringify(session));
export const clearSession = () => sessionStorage.removeItem(key);

export function backendUrl() {
  const url = new URL(window.location.href);
  url.port = '8080';
  url.pathname = '';
  url.search = '';
  url.hash = '';
  return url.origin;
}

export async function api<T = unknown>(path: string, init: RequestInit = {}): Promise<T> {
  const session = getSession();
  const response = await fetch(backendUrl() + path, {
    ...init,
    headers: {
      'Content-Type': 'application/json',
      ...(session ? {Authorization: `Bearer ${session.accessToken}`} : {}),
      ...init.headers,
    },
  });
  const text = await response.text();
  const data = text ? JSON.parse(text) : {};
  if (!response.ok) {
    throw new Error(data.error?.message || 'No se pudo completar la solicitud');
  }
  return data as T;
}
