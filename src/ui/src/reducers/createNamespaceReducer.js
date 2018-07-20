import {
  CREATE_NAMESPACE,
  NAMESPACE_MODAL_SHOW,
  NAMESPACE_MODAL_HIDE,
  NAMESPACE_ADD_ERROR,
} from '../actions/createNamespaceActions';

const initialState = {
  show: false,
  errorMessage: '',
};
export default function createNamespaceReducer(state = initialState, { type, payload }) {
  switch (type) {
    case NAMESPACE_MODAL_SHOW:
    case NAMESPACE_MODAL_HIDE:
    case NAMESPACE_ADD_ERROR:
    case CREATE_NAMESPACE:
      return payload;
    default:
      return state;
  }
}
