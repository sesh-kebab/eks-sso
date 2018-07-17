import { AUTHENTICATE_USER, LOGOUT_USER } from '../actions/authenticateActions'

export default function authReducer(state = {}, { type, payload }) {
  switch (type) {
    case AUTHENTICATE_USER:
      return payload.user;
    case LOGOUT_USER:
      return payload.user;
    default:
      return state;
  }
}