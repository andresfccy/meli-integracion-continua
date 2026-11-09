import {api} from '../../auth/session';

export async function gql<T>(
  query: string,
  variables: Record<string, unknown> = {},
): Promise<T> {
  const response = await api<{errors?: {message: string}[]; data: T}>('/graphql', {
    method: 'POST',
    body: JSON.stringify({query, variables}),
  });
  if (response.errors?.length) {
    throw new Error(response.errors[0].message);
  }
  return response.data as T;
}
