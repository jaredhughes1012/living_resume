
import { AccountInput, AuthData, Credentials, Identity, IdentityInput } from '@types';
import { create } from 'zustand';
import { devtools, persist } from 'zustand/middleware';
import client from './api';

interface AccountState {
  identity?: Identity;
  authToken: string;

  initiateNewAccount: (input: AccountInput) => Promise<void>;
  createIdentity: (input: IdentityInput) => Promise<Identity>;
  login: (input: Credentials) => Promise<Identity>;
}

const useAccountStore = create<AccountState>()(devtools(persist((set) => ({
  authToken: '',

  initiateNewAccount: async (input: AccountInput) => {
    const debug = import.meta.env.MODE === 'development';
    const res = await client.post(`/api/iam/v1/accounts/initiate?debug=${debug}`, input);
    if (debug) {
      console.log("Debug activation response", res.data);
    }
  },

  createIdentity: async (input: IdentityInput) => {
    const res = await client.post<AuthData>('/api/iam/v1/accounts', input);

    set(state => ({ ...state, authToken: res.data.authToken, identity: res.data.identity }));
    return res.data.identity;
  },

  login: async (input: Credentials) => {
    const res = await client.post<AuthData>('/api/iam/v1/authenticate', input);

    set(state => ({ ...state, authToken: res.data.authToken, identity: res.data.identity }));
    return res.data.identity;
  }
}), { name: 'account' })));

export default useAccountStore;