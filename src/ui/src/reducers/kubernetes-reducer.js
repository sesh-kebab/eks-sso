import { LIST_NAMESPACES } from '../actions/kubernetes-actions'

const defaultState = {namespaces:[]}

export default function kubernetesReducer(state = defaultState, { type, payload }) {
  switch (type) {
    case LIST_NAMESPACES:
      return payload
    default:
      return state;
  }
}