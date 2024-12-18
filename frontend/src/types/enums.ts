export enum ClassroomRole {
    PROFESSOR = "PROFESSOR",
    TA = "TA",
    STUDENT = "STUDENT",
  }
    
  export enum ClassroomUserStatus {
    REQUESTED = "REQUESTED",
    ACTIVE = "ACTIVE",
    ORG_INVITED = "ORG_INVITED",
    NOT_IN_ORG = "NOT_IN_ORG",
  }
  
  export enum OrgRole {
    ADMIN = "ADMIN",
    MEMBER = "MEMBER",
  }
  
  export function toClassroom(classroomUser: IClassroomUser) {
    return {
      id: classroomUser.classroom_id,
      name: classroomUser.classroom_name,
      org_id: classroomUser.org_id,
      org_name: classroomUser.org_name,
      created_at: classroomUser.classroom_created_at,
    }
  }

  export enum StudentWorkState {
    NOT_ACCEPTED = "NOT_ACCEPTED",
    ACCEPTED = "ACCEPTED", 
    STARTED = "STARTED",
    SUBMITTED = "SUBMITTED",
    GRADING_ASSIGNED = "GRADING_ASSIGNED",
    GRADING_COMPLETED = "GRADING_COMPLETED",
    GRADE_PUBLISHED = "GRADE_PUBLISHED"
  }


