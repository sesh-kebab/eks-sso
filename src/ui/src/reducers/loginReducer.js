import { LOGIN_ERROR, LOGIN_START } from '../actions/loginActions';

export default function authReducer(state = {}, { type, payload }) {
  switch (type) {
    case LOGIN_START:
      return payload;
    case LOGIN_ERROR:
      return payload.error;
    default:
      return state;
  }
}
