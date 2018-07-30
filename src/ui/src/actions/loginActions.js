export const LOGIN_START = 'login:LOGIN_START';
export const LOGIN_ERROR = 'login:LOGIN_ERROR';

const loginError = message => ({
  type: LOGIN_ERROR,
  payload: {
    error: {
      message: 'invalid username or password',
      internalMessage: message,
    },
  },
});

const loginStart = () => ({
  type: LOGIN_START,
  payload: {},
});

export { loginError, loginStart };
