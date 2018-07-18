import { LOGIN_ERROR, LOGIN_START } from '../actions/loginActions';

export default function authReducer(state = {}, { type, payload }) {
  switch (type) {
    case LOGIN_START:
    case LOGIN_ERROR:
      return payload.error;
    default:
      return state;
  }
}
