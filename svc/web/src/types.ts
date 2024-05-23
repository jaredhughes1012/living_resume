
export interface ChatInput {
  message: string;
}

export interface Identity {
  accountId: string;
  email: string;
  firstName: string;
  lastName: string;
}

export interface AccountInput {
  email: string;
}

export interface Credentials {
  email: string;
  password: string;
}

export interface IdentityInput {
  activationCode: string;
  firstName: string;
  lastName: string;
  accountId: string;
  credentials: Credentials;
}

export interface AuthData {
  identity: Identity;
  authToken: string;
}