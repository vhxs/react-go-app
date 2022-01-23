import * as React from 'react';
import Button from '@mui/material/Button';
import axios from 'axios';

function QueryButton(props) {
  return (
    <Button
      variant="contained"
      color="primary"
      onClick={async () => {
        var res = await axios.get(`http://localhost:8090/getComments`)
        props.setFunction([...props.results, {
            timestamp: new Date().getTime(),
            queryResult: res.data,
        }])
    }}>
        Get Comments!
    </Button>
  );
};

function PostButton(props) {
    return (
        <Button
            variant="contained"
            color="primary"
            onClick={async () => {
                axios.post(`http://localhost:8090/makeComment`, {
                    comment: props.data.current.value,
                })
            }}>
            Add Comment!
        </Button>
    );
}

export {QueryButton, PostButton};