interface ITemplateRepo {
    template_repo_id: number;
    template_repo_owner: string;
    template_repo_name: string;
}

interface ITemplatesResponse {
    templates: ITemplateRepo[];
}