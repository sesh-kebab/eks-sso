import { LOGIN_START, LOGIN_ERROR } from '../actions/loginActions'

export default function authReducer(state = {}, { type, payload }) {
  switch (type) {
    case LOGIN_START:
    case LOGIN_ERROR:
      return payload.error;
    default:
      return state;
  }
}