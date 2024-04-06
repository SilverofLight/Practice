import logo from './logo.svg';
import './App.css';

const type = {
  backgroundColor: "lightblue",
  borderColor: "black",
  borderRadius: "5px",
  margin: "3px",
  padding: "10px",

}

const displays = {
  display: "grid",
  gridTemplateAreas: "one two three",
  gridTemplateColumns: "1fr 1fr 1fr",
}

function Card(props) {
  return(
    <div className='card' style={type}>
      <h2>{props.h2}</h2>
      <h3>{props.h3}</h3>
    </div>
  );
}

function App() {
  return(
    <div className='App'>
      <h1>This is a Text of Card</h1>
      <div style={displays}>
        <Card h2="First Card's h2" h3="First Card's h3"/>
        <Card h2="Second Card's h2" h3="Second Card's h3"/>
        <Card h2="Third Card's h2" h3="Third Card's h3"/>
      </div>
    </div>
  )
}

export default App;
