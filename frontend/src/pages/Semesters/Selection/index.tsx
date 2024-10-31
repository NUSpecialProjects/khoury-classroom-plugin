// import React, { useEffect, useState, useContext } from "react";
// import { useNavigate } from "react-router-dom";
// import "./styles.css";
// import { FaChevronRight, FaChevronDown } from "react-icons/fa";
// import {
//   Table,
//   TableRow,
//   TableCell,
//   TableDiv,
// } from "@/components/Table/index.tsx";
// import { SelectedSemesterContext } from "@/contexts/selectedClassroom";
// import { Link } from "react-router-dom";
// import { getUserSemesters } from "@/api/semesters";

// const SemesterSelection: React.FC = () => {
//   const [semestersByOrg, setSemestersByOrg] = useState<{
//     [key: number]: ISemester[];
//   }>({});
//   const [collapsed, setCollapsed] = useState<{ [key: number]: boolean }>({});
//   const [loading, setLoading] = useState(true);
//   const { setSelectedSemester } = useContext(SelectedSemesterContext);

//   const navigate = useNavigate();

//   useEffect(() => {
//     const fetchSemesters = async () => {
//       try {
//         const data: IUserSemestersResponse = await getUserSemesters();
//         const groupedSemesters: { [key: number]: ISemester[] } = {};
//         const initialCollapsedState: { [key: number]: boolean } = {};

//         data.active_semesters
//           .concat(data.inactive_semesters)
//           .forEach((semester) => {
//             if (!groupedSemesters[semester.org_id]) {
//               groupedSemesters[semester.org_id] = [];
//               initialCollapsedState[semester.org_id] = false; // Set default to uncollapsed
//             }
//             groupedSemesters[semester.org_id].push(semester);
//           });

//         setSemestersByOrg(groupedSemesters);
//         setCollapsed(initialCollapsedState);
//       } catch (error) {
//         console.error("Error fetching semesters:", error);
//       } finally {
//         setLoading(false);
//       }
//     };

//     void fetchSemesters();
//   }, []);

//   const handleSemesterSelect = (semester: ISemester) => {
//     setSelectedSemester(semester);
//     navigate(`/app/dashboard`);
//   };

//   const toggleCollapse = (orgId: number) => {
//     setCollapsed((prev) => ({ ...prev, [orgId]: !prev[orgId] }));
//   };

//   const hasSemesters = Object.keys(semestersByOrg).length > 0;

//   return (
//     <div className="Selection">
//       <h1 className="Selection__title">Your Classrooms</h1>
//       <div className="Selection__tableWrapper">
//         {loading ? (
//           <p>Loading...</p>
//         ) : hasSemesters ? (
//           <Table cols={2} primaryCol={1}>
//             {Object.keys(semestersByOrg).map((orgId) => (
//               <React.Fragment key={orgId}>
//                 <TableRow className="HeaderRow" style={{ borderTop: "none" }}>
//                   <TableCell></TableCell>
//                   <TableCell>Organization Name</TableCell>
//                 </TableRow>
//                 <TableRow
//                   className={`ChildRow ${!collapsed[Number(orgId)] ? "TableRow--expanded" : ""}`}
//                 >
//                   <TableCell onClick={() => toggleCollapse(Number(orgId))}>
//                     {collapsed[Number(orgId)] ? (
//                       <FaChevronRight />
//                     ) : (
//                       <FaChevronDown />
//                     )}
//                   </TableCell>
//                   <TableCell>{`${semestersByOrg[Number(orgId)][0].org_name}`}</TableCell>
//                 </TableRow>
//                 {!collapsed[Number(orgId)] && (
//                   <TableDiv>
//                     <Table
//                       cols={3}
//                       primaryCol={0}
//                       className="DetailsTable SubTable"
//                     >
//                       <TableRow
//                         className="HeaderRow"
//                         style={{ borderTop: "none" }}
//                       >
//                         <TableCell>Classroom Name</TableCell>
//                         <TableCell className="status">Status</TableCell>
//                         <TableCell></TableCell>
//                       </TableRow>
//                       {semestersByOrg[Number(orgId)].map((semester) => (
//                         <TableRow
//                           key={`${semester.org_id}-${semester.classroom_id}`}
//                         >
//                           <TableCell>{semester.classroom_name}</TableCell>
//                           <TableCell className="status">
//                             {semester.active ? "Active" : "Inactive"}
//                           </TableCell>
//                           <TableCell>
//                             <button
//                               className="Selection__actionButton"
//                               onClick={() => handleSemesterSelect(semester)}
//                             >
//                               View
//                             </button>
//                           </TableCell>
//                         </TableRow>
//                       ))}
//                     </Table>
//                   </TableDiv>
//                 )}
//               </React.Fragment>
//             ))}
//           </Table>
//         ) : (
//           <div>
//             <p>You have no classes.</p>
//           </div>
//         )}
//       </div>
//       <div className="Selection__linkWrapper">
//         <Link to="/app/classroom/create">
//           {" "}
//           Create a new classroom instead →
//         </Link>
//       </div>
//     </div>
//   );
// };

// export default SemesterSelection;
