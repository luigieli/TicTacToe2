import React from "react";

const Square = React.memo(({ value, onClick }) => {
  return (
    <button className="square" onClick={onClick}>
      {value}
    </button>
  );
});

export default Square;
