import {
  CREATE_NAMESPACE,
  NAMESPACE_MODAL_SHOW,
} from '../actions/createNamespaceActions';

const initialState = {
  show: false,
  name: '',
};
export default function createNamespaceReducer(state = initialState, { type, payload }) {
  switch (type) {
    case NAMESPACE_MODAL_SHOW:
    case CREATE_NAMESPACE:
      return payload;
    default:
      return state;
  }
}
