export interface JobsResponse {
    totalCount: number;
    page:       number;
    pageSize:   number;
    totalPages: number;
    hasPrev:    boolean;
    hasNext:    boolean;
    guitars:    Job[];
}

export interface Job {
  id:               number;
  url:              string;
  title:            string;
  location:         string;
  minSalaryAtMonth: number;
  maxSalaryAtMonth: number;
  employmentType:   string;
  workPlace:        string;
  sourceSite:       string;
  createdAt:        string;
  updatedAt:        string;
  featureNamesCSV:  string;
  optionsCSV:       string;
  featureNames:     string[];
  options:          string[];
}