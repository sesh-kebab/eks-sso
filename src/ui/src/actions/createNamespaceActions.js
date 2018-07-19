import { getNamespaces } from './kubernetes-actions';
// import { onAddError } from './iamAddUserModalActions';

export const NAMESPACE_MODAL_SHOW = 'namespace-modal:show';
export const NAMESPACE_MODAL_HIDE = 'namespace-modal:hide';
export const CREATE_NAMESPACE = 'namespace:create';

const showNamespaceModal = () => ({
  type: NAMESPACE_MODAL_SHOW,
  payload: {
    show: true,
    name: '',
  },
});

const hideNamespaceModal = () => ({
  type: NAMESPACE_MODAL_HIDE,
  payload: {
    show: false,
    name: '',
  },
});

const createNamespace = (name) => (dispatch, getState) => {
  postData('/namespace?name=' + name).then(o => {
    return o.json();
  })
    .then(o => {
      dispatch(getNamespaces());
      return o.json();
    })
    .catch(e => {
      console.log(e.message);
      
      // dispatch(onAddError(e.message));
    });
};


const postData = (url) => {
  return fetch(url, {
    headers: {
      'Content-Type': 'application/json; charset=utf-8',
    },
    cache: 'no-cache',
    method: 'POST',
    credentials: 'same-origin',
    body: JSON.stringify({}),
  })
}

export { createNamespace, showNamespaceModal, hideNamespaceModal };


