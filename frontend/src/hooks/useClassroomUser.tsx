import { useState, useEffect } from "react";
import { getCurrentClassroomUser } from "@/api/classrooms";
import { ClassroomRole, requireAtLeastClassroomRole } from "@/types/users";
import { useNavigate } from "react-router-dom";

export function useClassroomUser(classroomId?: number, requiredRole?: ClassroomRole, redirectPath?: string) {
  const [classroomUser, setClassroomUser] = useState<IClassroomUser | null>(
    null
  );
  const [error, setError] = useState<Error | null>(null);
  const [loading, setLoading] = useState(true);

  const navigate = useNavigate();

  useEffect(() => {
    const fetchClassroomUser = async () => {
      if (classroomId) {
        await getCurrentClassroomUser(classroomId)
          .then((user) => {
            if (user.classroom_id === classroomId) {
              setClassroomUser(user);
              setError(null);
              if (requiredRole && !requireAtLeastClassroomRole(user.classroom_role, requiredRole)) {
                if (redirectPath) {
                  navigate(redirectPath);
                }
              }
            } else {
              setError(new Error("User is not in the specified classroom"));
              setClassroomUser(null);
              if (redirectPath) {
                navigate(redirectPath);
              }
            }
          })
          .catch((err) => {
            setError(err);
            setClassroomUser(null);
          })
          .finally(() => {
            setLoading(false);
          });
      } else {
        setClassroomUser(null);
        setError(null);
        setLoading(false);
      }
    };

    fetchClassroomUser();
  }, [classroomId]);

  return { classroomUser, error, loading };
}
