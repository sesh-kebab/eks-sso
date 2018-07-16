import { addStart, addSuccess, addError} from './iam-modal-actions'

export const ADD_CREDENTIALS = 'credentials:authenticate';
export const DELETE_CREDENTIALS = 'credentials:logout';
export const INVALID_CREDENTIALS = 'credentials:invalid';


export function addCredentials(accessId, secretKey) {
  return dispatch => {
    dispatch(addStart());
    const data = {
      accessId,
      secretKey
    };

    postData('/credentials', data).then(o => {
      return o.json();
    })
      .then(j => {
        dispatch(addSuccess());
        dispatch(add(j.username));
      })
      .catch(e => {
        dispatch(addError(e.message))
      })
  }
}

const add = (user) => {
  return {
    type: ADD_CREDENTIALS,
    payload: {
        user,
        valid: true,
    }
  }
}


const postData = (url, data) => {
  return fetch(url, {
    body: JSON.stringify(data),
    headers: {
      'Content-Type': 'application/json; charset=utf-8',
    },
    cache: 'no-cache',
    method: 'POST',
    credentials: 'same-origin',
  })
}