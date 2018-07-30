import { loginError, loginStart } from './loginActions';

export const AUTHENTICATE_USER = 'user:authenticate';
export const LOGOUT_USER = 'user:logout';

const onUserAuthenticated = (username, givenName, pictureUrl, clusterName) => ({
  type: AUTHENTICATE_USER,
  payload: {
    user: {
      name: username,
      givenName: givenName,
      pictureUrl: pictureUrl,
      authenticated: true,
      clusterName: clusterName,
    },
  },
});

const onUserLoggedOut = () => ({
  type: LOGOUT_USER,
  payload: {
    user: {
      name: 'anonymous',
      givenName: '',
      pictureUrl: '',
      authenticated: false,
      clusterName: '',
    },
  },
});

const authenticateUser = (username, password) => (dispatch, getState, api) => {
  const data = { username, password };

  dispatch(loginStart());

  api
    .postData('/authenticate', data)
    .then(resp => {
      if (!resp.ok) {
        throw Error(resp.statusText);
      }

      return resp.json();
    })
    .then(j => {
      dispatch(
        onUserAuthenticated(j.username, j.givenName, j.pictureUrl, j.clusterName, j.authenticated)
      );
    })
    .catch(e => {
      dispatch(loginError(e.message));
    });
};

//todo: make a server side request as well to update cookie
const logoutUser = () => dispatch => {
  dispatch(onUserLoggedOut());
};

export { authenticateUser, logoutUser };
