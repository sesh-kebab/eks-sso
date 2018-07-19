import { compose, lifecycle } from 'recompose';

import api from './../../api';
import KubeConfigComponent from './KubeConfigComponent';

const enhance = compose(
  // todo: add logic to show a loading indicator?
  lifecycle({
    state: { loading: true },
    componentDidMount() {
      api
        .getData('/cluster')
        .then(o => {
          return o.json();
        })
        .then(j => {
          this.setState({
            loading: false,
            kubeconfig: j.kubeconfig,
          });
        })
        .catch(e => {
          console.error(e);
        });
    },
  })
);

export default enhance(KubeConfigComponent);
