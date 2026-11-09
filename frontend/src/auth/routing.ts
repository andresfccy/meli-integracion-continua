export function defaultPath(roles: string[]) {
  if (roles.includes('seller')) return '/seller/orders';
  if (roles.includes('operations')) return '/operations/questions';
  if (roles.includes('buyer')) return '/buyer/orders';
  return '/login';
}
