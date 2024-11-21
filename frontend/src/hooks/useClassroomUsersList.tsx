import { useState, useEffect } from "react";
import { getClassroomUsers } from "@/api/classrooms";

export function useClassroomUsersList(classroomId?: number) {
  const [classroomUsers, setClassroomUsers] = useState<IClassroomUser[]>([]);
  const [error, setError] = useState<Error | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchClassroomUsers = async () => {
      if (classroomId) {
        await getClassroomUsers(classroomId)
          .then((users) => {
            setClassroomUsers(users);
            setError(null);
          })
          .catch((_) => {
            setError(new Error("Failed to fetch classroom users"));
            setClassroomUsers([]);
          })
          .finally(() => {
            setLoading(false);
          });
      } else {
        setClassroomUsers([]);
        setError(null);
        setLoading(false);
      }
    };

    fetchClassroomUsers();
  }, [classroomId]);

  return { classroomUsers, error, loading };
}
