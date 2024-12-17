export enum ClassroomRole {
  PROFESSOR = "PROFESSOR",
  TA = "TA",
  STUDENT = "STUDENT",
}

const roleValues = {
  [ClassroomRole.PROFESSOR]: 3,
  [ClassroomRole.TA]: 2,
  [ClassroomRole.STUDENT]: 1
};

export function requireAtLeastClassroomRole(role: ClassroomRole, requiredRole: ClassroomRole): boolean {
  return roleValues[role] >= roleValues[requiredRole];
}

export function requireGreaterClassroomRole(role: ClassroomRole, requiredRole: ClassroomRole): boolean {
  return roleValues[role] > roleValues[requiredRole];
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
