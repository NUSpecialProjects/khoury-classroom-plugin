/* eslint-disable */

import "./styles.css";
import { Classroom } from "@/types/classroom";

interface Props {
  availableClassrooms: Classroom[];
  unavailableClassrooms: Classroom[];
  selectedClassroom: Classroom | null;
  loading: Boolean;
  onSelect: (classroom: Classroom) => Promise<void>;
}

const ClassroomDropdown: React.FC<Props> = ({
  availableClassrooms,
  unavailableClassrooms,
  selectedClassroom,
  loading,
  onSelect,
}) => {
  return (
    <div className="form-group">
      <label htmlFor="classroom">Select Classroom:</label>
      <select
        id="classroom"
        value={selectedClassroom ? selectedClassroom.id : ""}
        onChange={async (e) => {
          const selectedId = Number(e.target.value);
          if (selectedId === -1) {
            window.open(
              "https://classroom.github.com/classrooms/new",
              "_blank"
            );
          } else {
            const selected = [
              ...availableClassrooms,
              ...unavailableClassrooms,
            ].find((classroom) => classroom.id === selectedId);
            if (selected) {
              await onSelect(selected);
            }
          }
        }}
      >
        {loading ? (
          <option value="" disabled>
            Loading classrooms...
          </option>
        ) : (
          <>
            <option value="" disabled>
              Select a classroom
            </option>
            {availableClassrooms.length > 0 && (
              <optgroup label="Available Classrooms">
                {availableClassrooms.map((classroom) => (
                  <option
                    key={classroom.id}
                    value={classroom.id}
                    title="This classroom is available to use"
                  >
                    {classroom.name} ✔️
                  </option>
                ))}
              </optgroup>
            )}
            {unavailableClassrooms.length > 0 && (
              <optgroup label="Unavailable Classrooms">
                {unavailableClassrooms.map((classroom) => (
                  <option
                    key={classroom.id}
                    value={classroom.id}
                    title="This classroom has already been used to create a semester"
                  >
                    {classroom.name} ❌
                  </option>
                ))}
              </optgroup>
            )}
          </>
        )}
        <option value="-1">Create New Classroom ➕</option>
      </select>
    </div>
  );
};

export default ClassroomDropdown;
