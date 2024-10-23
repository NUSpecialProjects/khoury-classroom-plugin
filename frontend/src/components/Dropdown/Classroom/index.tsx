import React from "react";
import '../styles.css'

interface Props {
  availableClassrooms: IClassroom[];
  unavailableClassrooms: IClassroom[];
  selectedClassroom: IClassroom | null;
  loading: boolean;
  onSelect: (classroom: IClassroom) => Promise<void>;
}

const ClassroomDropdown: React.FC<Props> = ({
  availableClassrooms,
  unavailableClassrooms,
  selectedClassroom,
  loading,
  onSelect,
}) => {
  return (
    <div className="Dropdown__wrapper">
      <label className="Dropdown__label" htmlFor="classroom">Select Classroom:</label>
      <select
        id="classroom"
        className="Dropdown"
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
          <option className="Dropdown__option" value="" disabled>
            Loading classrooms...
          </option>
        ) : (
          <>
            <option className="Dropdown__option" value="" disabled>
              Select a classroom
            </option>
            {availableClassrooms.length > 0 && (
              <optgroup label="Available Classrooms">
                {availableClassrooms.map((classroom) => (
                  <option
                    className="Dropdown__option"
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
                    className="Dropdown__option"
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
        <option className="Dropdown__option" value="-1">Create New Classroom ➕</option>
      </select>
    </div>
  );
};

export default ClassroomDropdown;
