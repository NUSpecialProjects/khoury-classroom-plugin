import React, { useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import "./styles.css";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { Link } from "react-router-dom";
import { getClassroomsInOrg } from "@/api/classrooms";
import useUrlParameter from "@/hooks/useUrlParameter";

const ClassroomSelection: React.FC = () => {
  const [classrooms, setClassrooms] = useState<IClassroom[]>([]);
  const orgID = useUrlParameter("org_id");
  const [loading, setLoading] = useState(true);
  const { setSelectedClassroom } = useContext(SelectedClassroomContext);

  const navigate = useNavigate();

  useEffect(() => {
    const fetchClassrooms = async () => {
      if (!orgID) {
        return;
      }
      setLoading(true);
      try {
        const org_id = parseInt(orgID);
        const data: IClassroomListResponse = await getClassroomsInOrg(org_id);
        const classrooms: IClassroom[] = data.classrooms;
        setClassrooms(classrooms);
      } catch (error) {
        console.error("Error fetching organizations and classrooms:", error);
      } finally {
        setLoading(false);
      }
    };

    void fetchClassrooms();
  }, [orgID]);

  const handleClassroomSelect = (classroom: IClassroom) => {
    setSelectedClassroom(classroom);
    navigate(`/app/dashboard`);
  };

  const hasClassrooms = classrooms.length > 0;

  return (
    <div className="Selection">
      <h1 className="Selection__title">Your Classrooms</h1>
      <div className="Selection__tableWrapper">
        {loading ? (
          <p>Loading...</p>
        ) : hasClassrooms ? (
          <div>
            {classrooms.map((classroom) => (
              <div
                key={classroom.id}
                className="Selection__tableRow"
                onClick={() => handleClassroomSelect(classroom)}
              >
                <p>{classroom.name}</p>
              </div>
            ))}
          </div>
        ) : (
          <div>
            <p>You have no classes.</p>
          </div>
        )}
      </div>
      <div className="Selection__linkWrapper">
        <Link to={`/app/classroom/create?org_id=${orgID}`}>
          {" "}
          Create a new classroom instead →
        </Link>
      </div>
      <div className="Selection__linkWrapper">
        <Link to={`/app/organization/select`}>
          {" "}
          Choose a different organization
        </Link>
      </div>
    </div>
  );
};

export default ClassroomSelection;
