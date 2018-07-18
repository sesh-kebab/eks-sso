import './index.css';
import { Provider } from 'react-redux';
import { HashRouter as Router } from 'react-router-dom';

import React from 'react';
import App from './App';
import configureStore from './configureStore'
import ReactDOM from 'react-dom';

import registerServiceWorker from './registerServiceWorker';

const store = configureStore();

ReactDOM.render(
  <Provider store={store}>
    <Router>
      <App />
    </Router>
  </Provider>,
  document.getElementById('root')
);
registerServiceWorker();
