export const IAM_MODAL_SHOW = 'iam-modal:show';
export const IAM_MODAL_HIDE = 'iam-modal:hide';
export const IAM_ADD_CREDENTIALS = 'iam-modal:add-credentials';
export const IAM_ADD_START = 'iam-modal:add-start';
export const IAM_ADD_ERROR = 'iam-modal:add-error';
export const IAM_ADD_SUCCESS = 'iam-modal:add-success';

const showIAMAddModal = () => ({
  type: IAM_MODAL_SHOW,
  payload: {
    show: true,
    message: '',
  },
});

const hideIAMAddModal = () => ({
  type: IAM_MODAL_SHOW,
  payload: {
    show: false,
    message: '',
  },
});

const onAddStart = () => ({
  type: IAM_ADD_START,
  payload: {
    show: true,
    message: '',
  },
});

const onAddError = message => ({
  type: IAM_ADD_ERROR,
  payload: {
    show: true,
    message: 'invalid access id and/or secret key',
    internalMessage: message,
  },
});

const onAddSuccess = () => ({
  type: IAM_ADD_SUCCESS,
  payload: {
    show: false,
    message: '',
  },
});

export { showIAMAddModal, hideIAMAddModal, onAddStart, onAddError, onAddSuccess };
