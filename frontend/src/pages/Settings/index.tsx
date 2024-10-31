import { activateSemester, deactivateSemester } from "@/api/semesters";
import ErrorMessage from "@/components/Error";
import { SelectedSemesterContext } from "@/contexts/selectedClassroom";
import { useContext, useState } from "react";

const Settings: React.FC = () => {
  const [error, setError] = useState<string | null>(null);
  const { selectedClassroom: selectedSemester, setSelectedSemester } = useContext(
    SelectedSemesterContext
  );

  const handleActivateClick = async () => {
    if (selectedSemester) {
      try {
        const newSemester = await activateSemester(
          selectedSemester.classroom_id
        );
        setSelectedSemester(newSemester);
        setError(null);
      } catch (err) {
        console.log(err);
        setError("Failed to activate the class. Please try again.");
      }
    }
  };

  const handleDeactivateClick = async () => {
    if (selectedSemester) {
      try {
        const newSemester = await deactivateSemester(
          selectedSemester.classroom_id
        );
        setSelectedSemester(newSemester);
        setError(null);
      } catch (err) {
        console.log(err);
        setError("Failed to deactivate the class. Please try again.");
      }
    }
  };

  return (
    <div>
      <h1>Settings</h1>
      <div>
        {selectedSemester && !selectedSemester.active && (
          <button onClick={handleActivateClick}>Activate Class</button>
        )}
        {selectedSemester && selectedSemester.active && (
          <button onClick={handleDeactivateClick}>Deactivate Class</button>
        )}
      </div>
      {error && <ErrorMessage message={error} />}
    </div>
  );
};

export default Settings;
