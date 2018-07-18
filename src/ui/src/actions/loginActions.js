export const LOGIN_START = 'login:LOGIN_START';
export const LOGIN_ERROR = 'login:LOGIN_ERROR';

const loginError = message => {
  return {
    type: LOGIN_ERROR,
    payload: {
      error: {
        message: 'invalid username or password',
        internalMessage: message,
      },
    },
  };
};

const loginStart = () => {
  return {
    type: LOGIN_ERROR,
    payload: {
      error: {
        message: '',
        internalMessage: '',
      },
    },
  };
};

export { loginError, loginStart };
