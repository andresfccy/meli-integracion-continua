import {describe, expect, it} from 'vitest';
import {defaultPath} from './routing';

describe('defaultPath', () => {
  it('applies the documented seller, operations, buyer priority', () => {
    expect(defaultPath(['buyer', 'operations', 'seller'])).toBe('/seller/orders');
    expect(defaultPath(['buyer', 'operations'])).toBe('/operations/questions');
    expect(defaultPath(['buyer'])).toBe('/buyer/orders');
  });

  it('returns login when no supported role is present', () => {
    expect(defaultPath([])).toBe('/login');
  });
});
