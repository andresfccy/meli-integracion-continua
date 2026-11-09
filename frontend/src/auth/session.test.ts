import {afterEach,beforeEach,describe,expect,it,vi} from 'vitest';
import {api,backendUrl,clearSession,getSession,setSession} from './session';

describe('browser session auth',()=>{
  beforeEach(()=>{clearSession();vi.unstubAllEnvs();});
  afterEach(()=>{vi.unstubAllGlobals();vi.unstubAllEnvs();});
  it('uses the web origin only in same-origin mode',()=>{
    vi.stubEnv('VITE_API_MODE','same-origin');
    expect(backendUrl()).toBe('http://localhost:3000');
  });
  it('preserves path and authorization through the same-origin proxy',async()=>{
    vi.stubEnv('VITE_API_MODE','same-origin');
    setSession({user:{id:'u',email:'a@b.co',displayName:'Ana',roles:['buyer']},accessToken:'signed-token'});
    const fetchMock=vi.fn().mockResolvedValue({ok:true,text:async()=>'{}'});
    vi.stubGlobal('fetch',fetchMock);
    await api('/graphql');
    expect(fetchMock).toHaveBeenCalledWith('http://localhost:3000/graphql',expect.objectContaining({headers:expect.objectContaining({Authorization:'Bearer signed-token'})}));
  });
  it('persists authentication only in sessionStorage',()=>{
    const session={user:{id:'u',email:'a@b.co',displayName:'Ana',roles:['buyer']},accessToken:'token'};
    setSession(session);
    expect(getSession()).toEqual(session);
  });
  it('derives backend port 8080 from the browser URL',()=>{
    expect(backendUrl()).toBe('http://localhost:8080');
  });
  it('adds the current access token to API requests',async()=>{
    setSession({user:{id:'u',email:'a@b.co',displayName:'Ana',roles:['buyer']},accessToken:'signed-token'});
    const fetchMock=vi.fn().mockResolvedValue({ok:true,text:async()=>JSON.stringify({value:1})});
    vi.stubGlobal('fetch',fetchMock);
    await api('/api/example');
    expect(fetchMock).toHaveBeenCalledWith(
      'http://localhost:8080/api/example',
      expect.objectContaining({headers:expect.objectContaining({Authorization:'Bearer signed-token'})}),
    );
    vi.unstubAllGlobals();
  });
});
