// Import useState, a React Hook that lets components remember values
import { useState } from "react";

// Define our main React component
function App() {

  // Create a state variable called showForm, initially false
  // setShowForm is the function used to update showForm
  const [showForm, setShowForm] = useState(false);

  // Stores the workout name netered by the user
  const [workoutName, setWorkoutName] = useState("");

  // Stores the last submitted workout name
  const [savedWorkout, setSavedWorkout] = useState("");


  function handleSubmit(e: React.SubmitEvent<HTMLFormElement>) {
    e.preventDefault(); // Prevents the page from refreshing

    setSavedWorkout(workoutName); // Store the submitted name
    setWorkoutName("");           // Clear the input
    setShowForm(false);           // Hide the form

    console.log(workoutName); // Prints the workout name
  }


  // Return the UI that React will display in the browser
  return (
    <div>

      {/* Display the saved workout if one exists */}
      {savedWorkout && <p>Workout: {savedWorkout}</p>}

      {/* Display the page heading */}
      <h1>Workout Tracker</h1>

      {/* When clicked, change showForm to true */}
      <button onClick={() => setShowForm(true)}>
        Add Workout
      </button>

      {/* Only display the form when showForm is true */}
      {showForm && (
        <form onSubmit={handleSubmit}>
          {/* Text field for entering the workout name */}
          <input 
            // Specifies that this is a text input
            type="text" 

            // Displays this hint when the input is empty
            placeholder="Workout name" 

            // Sets the input's displayed text to the workout Name state
            value={workoutName}

            // Runs wheneve the user types or changes the input
            onChange={(e) => setWorkoutName(e.target.value)}
            
            />

          {/* Button that submits the form */}
          <button type="submit">Save Workout</button>
        </form>
      )}
    </div>
  );
}

// Allow other files to import and use this component
export default App;
