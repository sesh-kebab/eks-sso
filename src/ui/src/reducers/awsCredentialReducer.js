import {
  ADD_CREDENTIALS,
  DELETE_CREDENTIALS,
  INVALID_CREDENTIALS,
} from '../actions/awsCredentialActions';

export default function credentialsReducer(state = { valid: false }, { type, payload }) {
  switch (type) {
    case ADD_CREDENTIALS:
      return payload;
    case DELETE_CREDENTIALS:
      return {};
    case INVALID_CREDENTIALS:
      return payload;
    default:
      return state;
  }
}
