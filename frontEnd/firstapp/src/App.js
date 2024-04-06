import logo from './logo.svg';
import './App.css';


function App() {
  function handleClick() {
    let computerNumber = Math.floor(Math.random() * 3) + 1;
    console.log(computerNumber);
    let userNumber = prompt("type a number");
    alert(`Computer number is ${computerNumber}, you guess ${userNumber}`);
  }

  return (
    <div>
      <h1>Add a button and handle a click event</h1>
      <button onClick={handleClick}>guess a number from 1~3</button>
    </div>
  )
}

export default App;
