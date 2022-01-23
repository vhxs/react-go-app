import './App.css';
import {QueryButton, PostButton} from './components/buttons';
import DenseTable from './components/queryTable';
import CommentBox from './components/commentBox';
import React, { useRef } from 'react';

function App() {
  const [queryResults, setQueryResults] = React.useState([{
    timestamp: 0,
    queryResult: "example result",
  }]);

  const fieldRef = useRef("");

  return (
    <div>
      <div className="App">
        <QueryButton results={queryResults} setFunction={setQueryResults}/>
        <PostButton data={fieldRef}/>
        <CommentBox inputRef={fieldRef}/>
      </div>
      <div>
        <DenseTable results={queryResults}/>
      </div>
    </div>
  );
}

export default App;
