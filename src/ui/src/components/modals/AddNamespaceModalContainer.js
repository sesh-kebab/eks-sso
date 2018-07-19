import { compose, withStateHandlers } from 'recompose';
import { connect } from 'react-redux';

import { createNamespace } from '../../actions/createNamespaceActions';
import { hideNamespaceModal } from '../../actions/createNamespaceActions';
import AddNamespaceModalComponent from './AddNamespaceModalComponent';

const mapStateToProps = state => ({
  namespaceModal: state.namespaceModal,
});

const mapActionsToProps = {
  createNamespace: createNamespace,
  hideModal: hideNamespaceModal,
};

const enhance = compose(
  connect(
    mapStateToProps,
    mapActionsToProps
  ),
  withStateHandlers(
    () => ({
      name: '',
    }),
    {
      onAccessNameChange: state => event => {
        state.name = event.target.value;
        return state;
      },
      onClose: (state, props) => () => {
        props.hideModal();
      },
      onSubmit: (state, props) => event => {
        event.preventDefault();
        props.createNamespace(state.name);
      },
    }
  )
);

export default enhance(AddNamespaceModalComponent);
