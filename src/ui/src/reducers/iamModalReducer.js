import { IAM_MODAL_SHOW, IAM_ADD_START, IAM_ADD_ERROR, IAM_ADD_SUCCESS } from '../actions/iamModalActions'

const initialState = {
  show: false,
  message: '',
}
export default function authReducer(state = initialState, { type, payload }) {
  switch (type) {
    case IAM_MODAL_SHOW:
    case IAM_ADD_START:
    case IAM_ADD_ERROR:
    case IAM_ADD_SUCCESS:
      return payload;
    default:
      return state;
  }
}