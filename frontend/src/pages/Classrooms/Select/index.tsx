import React, { useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import "./styles.css";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { Link } from "react-router-dom";
import { getClassroomsInOrg } from "@/api/classrooms";
import useUrlParameter from "@/hooks/useUrlParameter";
import { Table, TableRow, TableCell } from "@/components/Table";
import { MdAdd } from "react-icons/md";

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
        if (data.classrooms) {
          //TODO: this is broken
          setClassrooms(data.classrooms);
        }
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
          <Table cols={1}>
            <TableRow>
              <TableCell>Classroom Name</TableCell>
            </TableRow>
            {classrooms.map((classroom, i: number) => (
              <TableRow key={i} className="Selection__tableRow">
                <TableCell>
                  <div key={classroom.id} onClick={() => handleClassroomSelect(classroom)}>{classroom.name}</div>
                </TableCell>
              </TableRow>
            ))}
            <TableRow className="add-row">
              <TableCell style={{ textDecoration: "none" }}>
                <Link to={`/app/classroom/create?org_id=${orgID}`}>
                  <MdAdd />Create a new classroom
                </Link>
              </TableCell>
            </TableRow>
          </Table>
        ) : (
          <div>
            <p>You have no classes.</p>
          </div>
        )}
      </div>
  
      <div className="Selection__linkWrapper">
        <Link to={`/app/organization/select`}>
          {" "}
          Choose a different organization →
        </Link>
      </div>
    </div>
  );
};

export default ClassroomSelection;
