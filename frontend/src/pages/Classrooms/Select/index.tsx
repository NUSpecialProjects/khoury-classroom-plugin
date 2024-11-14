import React, { useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import "./styles.css";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { Link } from "react-router-dom";
import { getClassroomsInOrg } from "@/api/classrooms";
import useUrlParameter from "@/hooks/useUrlParameter";
import { Table, TableRow, TableCell } from "@/components/Table";
import EmptyDataBanner from "@/components/EmptyDataBanner";
import Button from "@/components/Button";
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
      } catch (_) {
        // do nothing
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
        {/* If the screen is loading, display a message, else render table */}
        {loading ? (
          <p>Loading...</p>
        ) : (
          <Table cols={1}>
            <TableRow>
              <TableCell>
                <div className="Selection__tableHeaderText">Classroom Name</div>
                <div className="Selection__tableHeaderButton">
                  <Button size="small" href={`/app/classroom/create?org_id=${orgID}`}>
                    <MdAdd /> New Classroom
                  </Button>
                </div>
              </TableCell>
            </TableRow>
            {/* If the org has classrooms, populate table, else display a message 
            TODO make alert for no classes*/}
            {hasClassrooms ? (
              classrooms.map((classroom, i) => (
                <TableRow key={i} className="Selection__tableRow">
                  <TableCell>
                    <div key={classroom.id} onClick={() => handleClassroomSelect(classroom)}>
                      {classroom.name}
                    </div>
                  </TableCell>
                </TableRow>
              ))
            ) : (
              <EmptyDataBanner>
                <div className="emptyDataBannerMessage">
                  You have no classes in this organization.
                  <br></br>
                  Please create a new classroom to get started.
                </div>
                <Button variant="secondary" href={`/app/classroom/create?org_id=${orgID}`}>
                    <MdAdd /> New Classroom
                  </Button>
              </EmptyDataBanner>
            )}
          </Table>
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
