import { API_URL } from "@/constants";
import axios from "axios";

axios.defaults.baseURL = API_URL;
axios.defaults.headers.common["Accept"] = "application/json";
axios.defaults.headers.post["Content-Type"] = "application/json";
axios.defaults.withCredentials = true;
