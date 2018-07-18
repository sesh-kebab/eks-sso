import { onAddError, onAddStart, onAddSuccess } from './iamAddUserModalActions';

export const ADD_CREDENTIALS = 'credentials:authenticate';
export const DELETE_CREDENTIALS = 'credentials:logout';
export const INVALID_CREDENTIALS = 'credentials:invalid';

const addCredentials = (accessId, secretKey) => (dispatch, getState, api) => {
  dispatch(onAddStart());
  const data = {
    accessId,
    secretKey,
  };

  api
    .postData('/credentials', data)
    .then(o => {
      return o.json();
    })
    .then(j => {
      dispatch(onAddSuccess());
      dispatch(add(j.username));
    })
    .catch(e => {
      dispatch(onAddError(e.message));
    });
};

const add = user => {
  return {
    type: ADD_CREDENTIALS,
    payload: {
      user,
      valid: true,
    },
  };
};

export { addCredentials };
