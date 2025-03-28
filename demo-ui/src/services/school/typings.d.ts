declare namespace API {
  interface classQuery  {
    class_name?: string;
    class_status?: number;
    school_uid?: number;
    course_id?: number;
    class_id?: number;
    client_class_id?: number;
    st_id?: number;
    ass_st_id?: number;
    time_str?: string;
    time_range?: [dayjs, dayjs],
  }
  type classItem = {
    class_btime?: string;
    class_etime?: string;
    class_id?: string;
    school_uid?: string;
    school_name?: string;
    course_id?: string;
    course_name?: string;
    class_name?: string;
    main_teacher?: string;
    assistants?: string;
    client_class_id?: string;
    seat_num?: string;
    folder_name?: string;
    folder_path?: string;
    record_state?: string;
    live_state?: string;
    open_state?: string;
    student_num?: string;
    student_list_url?: string;
    class_status?: string;
    class_status_text?: string;
    addtime?: string;
    st_id?: string;
    ass_st_id?: string;
    record_url?: string;
  }
  type classList = {
    data?: classItem[];
    /** 列表的内容总数 */
    error_info?: {[key: string]: any};
  }
}