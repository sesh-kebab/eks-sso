import { onAddError, onAddStart, onAddSuccess } from './iamAddUserModalActions';

export const ADD_CREDENTIALS = 'credentials:authenticate';
export const DELETE_CREDENTIALS = 'credentials:logout';
export const INVALID_CREDENTIALS = 'credentials:invalid';

const add = user => ({
  type: ADD_CREDENTIALS,
  payload: {
    user,
    valid: true,
  },
});

const addCredentials = (accessId, secretKey) => (dispatch, getState, api) => {
  dispatch(onAddStart());
  const data = {
    accessId,
    secretKey,
  };

  api
    .postData('/credentials', data)
    .then(resp => {
      if (!resp.ok) {
        throw Error(resp.statusText);
      }

      return resp.json();
    })
    .then(j => {
      dispatch(add(j.username));
      dispatch(onAddSuccess());
    })
    .catch(e => {
      dispatch(onAddError(e.message));
    });
};

export { addCredentials };
