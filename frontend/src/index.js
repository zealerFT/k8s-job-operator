import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import MyForm from './k8s-form';

ReactDOM.render(<MyForm />, document.getElementById('my-form-root'));

reportWebVitals();