using Newtonsoft.Json;
using System;
using System.IO;

namespace Example {
    public static class Example {
        public static void Main() {
            exampleJson();
            exampleProto();
        }

        /// <summary>
        /// json�d��
        /// </summary>
        private static void exampleJson() {
            // �n�ϥ�sheeter, �����إ��~�Ӧ�sheeterJson.Loader������Ū����
            var loader = new JsonFileLoader();
            // ���۫إ�sheeterJson.Depot����, �o�O�s������Ƴ̥D�n������
            // �n�O�o����إߪ�Ū�����]�w�i�h
            var depot = new SheeterJson.Depot() { Loader = loader };

            // �M�����FromData(�άOMergeData)�禡��Ū�������
            if (depot.FromData() == false) {
                Console.WriteLine("json failed: from data failed");
                return;
            }

            // ����N�i�H��Depot���U���U�Ӫ�檫��Ө��θ�Ƥ��e
            if (depot.ExampleData.TryGetValue(1, out var data)) {
                Console.WriteLine(JsonConvert.SerializeObject(data));
                Console.WriteLine("json success");
            } else {
                Console.WriteLine("json failed: get data failed");
            }
        }

        /// <summary>
        /// proto�d��
        /// </summary>
        private static void exampleProto() {
            // �n�ϥ�sheeter, �����إ��~�Ӧ�sheeterProto.Loader������Ū����
            var loader = new ProtoFileLoader();
            // ���۫إ�sheeterProto.Depot����, �o�O�s������Ƴ̥D�n������
            // �n�O�o����إߪ�Ū�����]�w�i�h
            var depot = new SheeterProto.Depot() { Loader = loader };

            // �M�����FromData(�άOMergeData)�禡��Ū�������
            if (depot.FromData() == false) {
                Console.WriteLine("proto failed: from data failed");
                return;
            }

            // ����N�i�H��Depot���U���U�Ӫ�檫��Ө��θ�Ƥ��e
            if (depot.ExampleData.TryGetValue(1, out var data)) {
                Console.WriteLine(data);
                Console.WriteLine("proto success");
            } else {
                Console.WriteLine("proto failed: get data failed");
            }
        }
    }

    /// <summary>
    /// json�ɮ�Ū����
    /// </summary>
    class JsonFileLoader : SheeterJson.Loader {
        /// <summary>
        /// �Ω�B�zŪ����ƿ��~, �d�Ҥ��u�O��¦L�X���~�T��
        /// </summary>
        /// <param name="name">�ɮצW��</param>
        /// <param name="message">���~�T��</param>
        public void Error(string name, string message) {
            Console.WriteLine(name + ": json file load failed: " + message);
        }

        /// <summary>
        /// �Ω�Ū������ɮ�, Depot�|���ѵ��A�ɮצW��, ���ɦW, ����W��
        /// �ϥΪ̻ݭn�̾a�H�W��T��Ū������ɮ�, �æ^�Ǹ�Ƶ�Depot
        /// </summary>
        /// <param name="name">�ɮצW��</param>
        /// <param name="ext">���ɦW</param>
        /// <param name="fullname">����W��</param>
        /// <returns></returns>
        public string Load(string name, string ext, string fullname) {
            // �]���u�@���|�b bin/Debug/net5.0/ ���U,�ҥH�u�n�[3�� ".."�F
            return File.ReadAllText(Path.Combine("..", "..", "..", "json", "data", fullname));
        }
    }

    /// <summary>
    /// proto�ɮ�Ū����
    /// </summary>
    class ProtoFileLoader : SheeterProto.Loader {
        /// <summary>
        /// �Ω�B�zŪ����ƿ��~, �d�Ҥ��u�O��¦L�X���~�T��
        /// </summary>
        /// <param name="name">�ɮצW��</param>
        /// <param name="message">���~�T��</param>
        public void Error(string name, string message) {
            Console.WriteLine(name + ": proto file load failed: " + message);
        }

        /// <summary>
        /// �Ω�Ū������ɮ�, Depot�|���ѵ��A�ɮצW��, ���ɦW, ����W��
        /// �ϥΪ̻ݭn�̾a�H�W��T��Ū������ɮ�, �æ^�Ǹ�Ƶ�Depot
        /// </summary>
        /// <param name="name">�ɮצW��</param>
        /// <param name="ext">���ɦW</param>
        /// <param name="fullname">����W��</param>
        /// <returns></returns>
        public byte[] Load(string name, string ext, string fullname) {
            // �]���u�@���|�b bin/Debug/net5.0/ ���U,�ҥH�u�n�[3�� ".."�F
            return File.ReadAllBytes(Path.Combine("..", "..", "..", "proto", "data", fullname));
        }
    }
}


