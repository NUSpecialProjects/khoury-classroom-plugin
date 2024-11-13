import { AssignmentFormData } from "@/components/MultiStepForm/Interfaces/CreateAssignment";

const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getAssignments = async (
  classroomId: number
): Promise<IAssignmentOutline[]> => {
  const result = await fetch(
    `${base_url}/classrooms/classroom/${classroomId}/assignments`,
    {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  if (!result.ok) {
    throw new Error("Network response was not ok");
  }

  const data = await result.json();

  return data.assignment_outlines as IAssignmentOutline[];
};

export const getAssignmentIndirectNav = async (
  classroomid: number,
  assignmentID: number
): Promise<IAssignmentOutline> => {
  const result = await fetch(`${base_url}/classrooms/classroom/${classroomid}/assignments/assignment/${assignmentID}`, {
    method: 'GET',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!result.ok) {
    throw new Error("Network response was not ok");
  }

  const data: IAssignmentOutline = (await result.json() as IAssignmentOutlineResponse).assignment_outline

  return data
};

export const createAssignmentTemplate = async (
  classroomId: number,
  assignment: IRepository
): Promise<IRepository> => {
  const result = await fetch(
    `${base_url}/classrooms/classroom/${classroomId}/assignments/template`,
    {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        template_repo_id: assignment.id,
        template_repo_owner: assignment.owner,
      })
    }
  );

  if (!result.ok) {
    throw new Error("Network response was not ok");
  }

  const data = (await result.json())

  return data.assignment_template as IRepository
};

export const acceptAssignment = async (orgName: string, repoName: string, classroomID: number, assignmentName: string) => {
  const result = await fetch(
      `${base_url}/classrooms/classroom/${classroomID}/assignments/accept`,
      {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          org_name: orgName,
          repo_name: repoName,
          assignment_name: assignmentName,
          assignment_id: 1,
          org_id: 182810684
        }),
      }
    );

    if (!result.ok) {
      throw new Error(result.statusText);
    }
};

export const createAssignment = async (
  templateId: number,
  assignment: AssignmentFormData
): Promise<AssignmentFormData> => {
  const result = await fetch(
    `${base_url}/classrooms/classroom/${assignment.classroomId}/assignments`,
    {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        template_id: templateId,
        name: assignment.assignmentName,
        classroom_id: assignment.classroomId,
        group_assignment: assignment.groupAssignment,
        main_due_date: assignment.mainDueDate,
      })
    }
  );

  if (!result.ok) {
    throw new Error("Network response was not ok");
  }

  const data = (await result.json())

  return data.assignment_outline as AssignmentFormData
};
