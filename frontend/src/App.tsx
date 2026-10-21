// Import useState, a React Hook that lets components remember values
import { useState } from "react";
import { DayPicker } from "react-day-picker";

// Import the day picker package for adding a calendar style date selection
import "react-day-picker/style.css";

// Define our main React component
function App() {

  // Create a state variable called showForm, initially false
  // setShowForm is the function used to update showForm
  const [showForm, setShowForm] = useState(false);

  // Stores the workout name netered by the user
  const [workoutName, setWorkoutName] = useState("");

  // Stores the selected date; defaults to today
  const [workoutDate, setWorkoutDate] = useState<Date | undefined>(
    new Date()
  )

  // Stores the number of sets and reps
  const [sets, setSets] = useState("");
  const [reps, setReps] = useState("");

  // Stores the last submitted workout name
  const [savedWorkout, setSavedWorkout] = useState("");

  // Send the workout to Go backend
  async function handleSubmit(e: React.SubmitEvent<HTMLFormElement>) {
    e.preventDefault();

    // Send POST request to Go
    const response = await fetch("/api/workouts", {
      method: "POST",
      headers: {
        "Content-Type":
          "application/json",
      },
      body: JSON.stringify({
        name: workoutName,
        date: workoutDate
          ? `${workoutDate.getFullYear()}-${String(workoutDate.getMonth() + 1).padStart(2, "0")}-${String(workoutDate.getDate()).padStart(2, "0")}`
          : null,
        sets: Number(sets),
        reps: Number(reps),
      }),
    });

    // Check whether Go successfully saved the workout
    if (!response.ok) {
      const errorMessage = await response.text();
      console.error("Failed to save workout:", errorMessage);
      return;
    }

    // Update the frontend after successful submission
    setSavedWorkout(workoutName);
    setWorkoutName("");
    setShowForm(false);
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

          {/*Interactive workout calendar */}
          <DayPicker
            mode="single"
            selected={workoutDate}
            onSelect={setWorkoutDate}
          />

          {/*Number of sets */}
          <input
            type="number"
            min="1"
            placeholder="Sets"
            value={sets}
            onChange={(e) => setSets(e.target.value)}
           />

          {/* Number of reps */}
          <input
            type="number"
            min="1"
            placeholder="Reps"
            value={reps}
            onChange={(e) => setReps(e.target.value)} 
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
