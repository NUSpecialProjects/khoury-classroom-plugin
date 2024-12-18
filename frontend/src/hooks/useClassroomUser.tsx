import { useQuery } from "@tanstack/react-query";
import { getCurrentClassroomUser } from "@/api/classrooms";

export function useClassroomUser(classroomId?: number) {
  const { data: classroomUser, error, isLoading } = useQuery({
    queryKey: ['classroomUser', classroomId],
    queryFn: async () => {
      if (!classroomId) return null;
      const user = await getCurrentClassroomUser(classroomId);
      if (user.classroom_id !== classroomId) {
        throw new Error("User is not in the specified classroom");
      }
      return user;
    },
    enabled: !!classroomId,
    retry: false
  });

  return { 
    classroomUser: classroomUser || null, 
    error: error as Error | null,
    loading: isLoading 
  };
}
