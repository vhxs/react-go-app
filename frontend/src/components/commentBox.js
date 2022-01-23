import * as React from 'react';
import TextField from '@mui/material/TextField';

export default function CommentBox(props) {
  return (
    <TextField
      id="standard-basic"
      label="Comment"
      variant="standard"
      inputRef={props.inputRef}
    />
  );
}