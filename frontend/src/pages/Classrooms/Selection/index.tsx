import React, { useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import "./styles.css";
import { FaChevronRight, FaChevronDown } from "react-icons/fa";
import {
  Table,
  TableRow,
  TableCell,
  TableDiv,
} from "@/components/Table/index.tsx";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { Link } from "react-router-dom";
import { getUserOrgsAndClassrooms } from "@/api/classrooms";
import { IOrganization, IClassroom } from "@/types";

const ClassroomSelection: React.FC = () => {
  const [classroomsByOrg, setClassroomsByOrg] = useState<
    Map<number, IClassroom[]>
  >(new Map());
  const [collapsed, setCollapsed] = useState<Map<number, boolean>>(new Map());
  const [loading, setLoading] = useState(true);
  const { setSelectedClassroom } = useContext(SelectedClassroomContext);

  const navigate = useNavigate();

  useEffect(() => {
    const fetchSemesters = async () => {
      try {
        const data: Map<IOrganization, IClassroom[]> = (
          await getUserOrgsAndClassrooms()
        ).orgs_and_classrooms;
        const initialCollapsedState: Map<number, boolean> = new Map();

        const transformedData: Map<number, IClassroom[]> = new Map();
        data.forEach((classrooms, org) => {
          transformedData.set(org.id, classrooms);
          initialCollapsedState.set(org.id, true);
        });

        setClassroomsByOrg(transformedData);
        setCollapsed(initialCollapsedState);
      } catch (error) {
        console.error("Error fetching organizations and classrooms:", error);
      } finally {
        setLoading(false);
      }
    };

    void fetchSemesters();
  }, []);

  const handleClassroomSelect = (classroom: IClassroom) => {
    setSelectedClassroom(classroom);
    navigate(`/app/dashboard`);
  };

  const toggleCollapse = (orgId: number) => {
    setCollapsed((prev) => {
      const newCollapsed = new Map(prev);
      newCollapsed.set(orgId, !prev.get(orgId));
      return newCollapsed;
    });
  };

  const hasSemesters = classroomsByOrg.size > 0;

  return (
    <div className="Selection">
      <h1 className="Selection__title">Your Classrooms</h1>
      <div className="Selection__tableWrapper">
        {loading ? (
          <p>Loading...</p>
        ) : hasSemesters ? (
          <Table cols={3} primaryCol={1}>
            <TableRow className="HeaderRow" style={{ borderTop: "none" }}>
              <TableCell></TableCell>
              <TableCell>Organization Name</TableCell>
              <TableCell>
                <a
                  href="https://github.com/organizations/plan"
                  className="Selection__actionButton"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  +
                </a>
              </TableCell>
            </TableRow>
            {Array.from(classroomsByOrg.entries()).map(
              ([orgId, classrooms]) => (
                <React.Fragment key={orgId}>
                  <TableRow
                    className={`ChildRow ${!collapsed.get(orgId) ? "TableRow--expanded" : ""}`}
                  >
                    <TableCell onClick={() => toggleCollapse(orgId)}>
                      {collapsed.get(orgId) ? (
                        <FaChevronRight />
                      ) : (
                        <FaChevronDown />
                      )}
                    </TableCell>
                    <TableCell>{`${classrooms[0].org_name}`}</TableCell>
                    <TableCell></TableCell>
                  </TableRow>
                  {!collapsed.get(orgId) && (
                    <TableDiv>
                      <Table
                        cols={2}
                        primaryCol={0}
                        className="DetailsTable SubTable"
                      >
                        <TableRow
                          className="HeaderRow"
                          style={{ borderTop: "none" }}
                        >
                          <TableCell>Classroom Name</TableCell>
                          <TableCell>
                            <a
                              href="https://classroom.github.com/classrooms/new" //TODO: this should link to OUR classroom page
                              className="Selection__actionButton"
                              target="_blank"
                              rel="noopener noreferrer"
                            >
                              +
                            </a>
                          </TableCell>
                        </TableRow>
                        {classrooms.map((classroom) => (
                          <TableRow key={classroom.id}>
                            <TableCell>{classroom.name}</TableCell>
                            <TableCell>
                              <button
                                className="Selection__actionButton"
                                onClick={() => handleClassroomSelect(classroom)}
                              >
                                View
                              </button>
                            </TableCell>
                          </TableRow>
                        ))}
                      </Table>
                    </TableDiv>
                  )}
                </React.Fragment>
              )
            )}
          </Table>
        ) : (
          <div>
            <p>You have no classes.</p>
          </div>
        )}
      </div>
      <div className="Selection__linkWrapper">
        <Link to="/app/classroom/create">
          {" "}
          Create a new classroom instead →
        </Link>
      </div>
    </div>
  );
};

export default ClassroomSelection;
