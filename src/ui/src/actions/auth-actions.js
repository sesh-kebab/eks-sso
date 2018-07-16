import {loginStart, loginError} from './login-actions';

export const AUTHENTICATE_USER = 'user:authenticate';
export const LOGOUT_USER = 'user:logout';

const updateAuthenticatedUser = (
  username, givenName, pictureUrl, clusterName
) => {
  return {
    type: AUTHENTICATE_USER,
    payload: {
      user: {
        name: username,
        givenName: givenName,
        pictureUrl: pictureUrl,
        authenticated: true,
        clusterName: clusterName,
      }
    }
  };
}

const logoutUser = () => {
  return {
    type: LOGOUT_USER,
    payload: {
      user: {
        name: 'anonymous',
        givenName: '',
        pictureUrl: '',
        authenticated: false,
        clusterName: '',
      }
    }
  }
}

export function authenticate(username, password) {
  return dispatch => {
    const data = {
      username: username,
      password: password
    };

    dispatch(loginStart());

    postData('/authenticate', data).then(o => {
      if (o.status >= 400) {
        throw Error(o.statusText);
      }

      return o.json();
    })
      .then(j => {
        dispatch(updateAuthenticatedUser(
          j.username,
          j.givenName,
          j.pictureUrl,
          j.clusterName,
          j.authenticated,
        ))
      })
      .catch(e => {
        dispatch(loginError(e.message))
      })
  }
}

export function logout() {
  return dispatch => {
    dispatch(logoutUser())
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
  });
}